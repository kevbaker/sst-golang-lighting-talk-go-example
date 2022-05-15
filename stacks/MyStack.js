import * as sst from "@serverless-stack/resources";

export default class MyStack extends sst.Stack {
  constructor(scope, id, props) {
    super(scope, id, props);

    // Create a HTTP API
    const api = new sst.Api(this, "Api", {
      defaultFunctionProps: {
        srcPath: 'backend/mainmodule',  // use path to application root
        environment: {
          HELLO_MESSAGE: "Hello World"
        },
      },
      routes: {
        "GET /": "cmd/handlers/exampleapp/main.go",          // use path relative to application root
        "GET /{count}/{samplerate}": "cmd/handlers/exampleapp/main.go",   // use path relative to application root
      }
    });

    // Show the endpoint in the output
    this.addOutputs({
      "ApiEndpoint": api.url,
    });
  }
}
