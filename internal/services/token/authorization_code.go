// Licensed to SolID under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. SolID licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package token

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/golang/protobuf/ptypes/wrappers"

	"go.zenithar.org/solid/pkg/rfcerrors"
	"go.zenithar.org/solid/pkg/storage"
	"go.zenithar.org/solid/pkg/types"

	corev1 "go.zenithar.org/solid/api/gen/go/oidc/core/v1"
	registrationv1 "go.zenithar.org/solid/api/gen/go/oidc/registration/v1"
)

func (s *service) authorizationCode(ctx context.Context, client *registrationv1.Client, req *corev1.TokenRequest) (*corev1.TokenResponse, error) {
	res := &corev1.TokenResponse{}
	grant := req.GetAuthorizationCode()

	// Check parameters
	if client == nil {
		res.Error = rfcerrors.ServerError("")
		return res, fmt.Errorf("unable to process with nil client")
	}
	if req == nil {
		res.Error = rfcerrors.ServerError("")
		return res, fmt.Errorf("unable to process with nil request")
	}
	if grant == nil {
		res.Error = rfcerrors.ServerError("")
		return res, fmt.Errorf("unable to process with nil grant")
	}

	// Validate client capabilities
	if !types.StringArray(client.GrantTypes).Contains(grantTypeAuthorizationCode) {
		res.Error = rfcerrors.UnsupportedGrantType("")
		return res, fmt.Errorf("client doesn't support 'authorization_code' as grant type")
	}

	// Validate request
	if grant.Code == "" || grant.CodeVerifier == "" || grant.RedirectUri == "" {
		res.Error = rfcerrors.InvalidGrant("")
		return res, fmt.Errorf("invalid authorization request: code, code_verifier and redirect_uri are mandatory")
	}

	// Retrieve authorization request from code
	ar, err := s.authorizationRequests.GetByCode(ctx, grant.Code)
	if err != nil {
		if err != storage.ErrNotFound {
			res.Error = rfcerrors.ServerError("")
		} else {
			res.Error = rfcerrors.InvalidGrant("")
		}
		return res, fmt.Errorf("unable to retrieve authorization request from code '%s': %w", grant.Code, err)
	}

	// Validate redirectUri
	if ar.RedirectUri != grant.RedirectUri {
		res.Error = rfcerrors.InvalidGrant(ar.State)
		return res, fmt.Errorf("invalid authorization request: request_uri from request '%s' and token '%s' must be identic", ar.RedirectUri, grant.RedirectUri)
	}
	if !types.StringArray(client.RedirectUris).Contains(grant.RedirectUri) {
		res.Error = rfcerrors.InvalidGrant(ar.State)
		return res, fmt.Errorf("invalid authorization request: request_uri from request '%s' and client '%s' must be validated", grant.RedirectUri, client.RedirectUris)
	}

	// Check PKCE verifier
	// https://www.rfc-editor.org/rfc/rfc7636.txt
	switch ar.CodeChallengeMethod {
	case "S256", "s256":
		h := sha256.Sum256([]byte(grant.CodeVerifier))
		computedVerifier := base64.RawURLEncoding.EncodeToString(h[:])
		if computedVerifier != ar.CodeChallenge {
			res.Error = rfcerrors.InvalidGrant(ar.State)
			return res, fmt.Errorf("unable to validate PKCE code_verifier `%s` and code_challenge `%s`", computedVerifier, ar.CodeChallenge)
		}
	default:
		res.Error = rfcerrors.InvalidGrant(ar.State)
		return res, fmt.Errorf("invalid code_challenge_method in request `%s`", ar.CodeChallengeMethod)
	}

	// Validate scopes
	scopes := types.StringArray(strings.Fields(ar.Scope))

	// Generate OpenID tokens (AT / RT / IDT)
	if scopes.Contains("openid") {
		// Create openid token holder
		var (
			err      error
			oidToken = &corev1.OpenIDToken{}
		)

		// Generate an access token
		oidToken.AccessToken, err = s.accessTokenGenerator.Generate(ctx)
		if err != nil {
			res.Error = rfcerrors.ServerError(ar.State)
			return res, fmt.Errorf("unbale to generate an accessToken: %w", err)
		}

		// Check for offline_access token, it means refresh token generation
		if scopes.Contains("offline_access") {
			// Generate an access token
			rt, err := s.accessTokenGenerator.Generate(ctx)
			if err != nil {
				res.Error = rfcerrors.ServerError(ar.State)
				return res, fmt.Errorf("unbale to generate an accessToken: %w", err)
			}

			// Assing refresh token
			oidToken.RefreshToken = &wrappers.StringValue{
				Value: rt,
			}
		}

		// Set "Bearer" token type
		oidToken.TokenType = "Bearer"
		oidToken.ExpiresIn = 3600

		// Assign openid tokens to result.
		res.Openid = oidToken
	}

	// No error
	return res, nil
}