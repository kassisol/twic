# HBM TWIC

TWIC is an open source project for managing Docker certificates to connect to the Docker daemon using TLS.

1. Add a TSA url and login using credentials.
4. TWIC request a certificate with the auto-generated private key for the profile using the token provided when authenticated to TSA.
2. If user authorized to use the Docker host, TSA sends a token to TWIC.
5. If authorized, CA sends the new certificate to TWIC.
3. User add a new profile to connect to Docker host.
6. User can use new profile to set Docker environment variables for connecting to Docker host using TLS.

## Getting Started & Documentation

All documentation is available on the [Harbormaster website](http://harbormaster.io/docs/twic/).

## User Feedback

### Issues

If you have any problems with or questions about this application, please contact us through a [GitHub](https://github.com/kassisol/twic/issues) issue.
