# Dockerfile validation

This **WIP project** focuses on automating the validation process of a Dockerfile by checking if the Dockerfile follows security best-practices.

- Uses [conftest](https://www.conftest.dev/) to parse the Dockerfile to get structured JSON data to be fed to Cue/OPA
    - `conftest parse Dockerfile > <Dockerfile JSON File NAME>` to convert Dockerfile to JSON
- `validate-input-dockerfile.rego` contains few policies written REGO to be used withh OPA
- `dockerfile-schema.cue` contains a Cue definition of a Dockerfile
- `test-dockerfile.json` contains a sample multi-stage Dockerfile in JSON format 
- [OPA Playground](https://play.openpolicyagent.org/p/OSMY88VouF) test link for the same policies
- `validate-dockerfile.rego` and `validate-input-dockerfile.rego` are validation policies written in Rego.

