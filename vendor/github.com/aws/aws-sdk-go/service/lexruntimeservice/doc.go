// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package lexruntimeservice provides the client and types for making API
// requests to Amazon Lex Runtime Service.
//
// Amazon Lex provides both build and runtime endpoints. Each endpoint provides
// a set of operations (API). Your conversational bot uses the runtime API to
// understand user utterances (user input text or voice). For example, suppose
// a user says "I want pizza", your bot sends this input to Amazon Lex using
// the runtime API. Amazon Lex recognizes that the user request is for the OrderPizza
// intent (one of the intents defined in the bot). Then Amazon Lex engages in
// user conversation on behalf of the bot to elicit required information (slot
// values, such as pizza size and crust type), and then performs fulfillment
// activity (that you configured when you created the bot). You use the build-time
// API to create and manage your Amazon Lex bot. For a list of build-time operations,
// see the build-time API, .
//
// See https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28 for more information on this service.
//
// See lexruntimeservice package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/lexruntimeservice/
//
// Using the Client
//
// To use the client for Amazon Lex Runtime Service you will first need
// to create a new instance of it.
//
// When creating a client for an AWS service you'll first need to have a Session
// already created. The Session provides configuration that can be shared
// between multiple service clients. Additional configuration can be applied to
// the Session and service's client when they are constructed. The aws package's
// Config type contains several fields such as Region for the AWS Region the
// client should make API requests too. The optional Config value can be provided
// as the variadic argument for Sessions and client creation.
//
// Once the service's client is created you can use it to make API requests the
// AWS service. These clients are safe to use concurrently.
//
//   // Create a session to share configuration, and load external configuration.
//   sess := session.Must(session.NewSession())
//
//   // Create the service's client with the session.
//   svc := lexruntimeservice.New(sess)
//
// See the SDK's documentation for more information on how to use service clients.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws package's Config type for more information on configuration options.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon Lex Runtime Service client LexRuntimeService for more
// information on creating the service's client.
// https://docs.aws.amazon.com/sdk-for-go/api/service/lexruntimeservice/#New
//
// Once the client is created you can make an API request to the service.
// Each API method takes a input parameter, and returns the service response
// and an error.
//
// The API method will document which error codes the service can be returned
// by the operation if the service models the API operation's errors. These
// errors will also be available as const strings prefixed with "ErrCode".
//
//   result, err := svc.PostContent(params)
//   if err != nil {
//       // Cast err to awserr.Error to handle specific error codes.
//       aerr, ok := err.(awserr.Error)
//       if ok && aerr.Code() == <error code to check for> {
//           // Specific error code handling
//       }
//       return err
//   }
//
//   fmt.Println("PostContent result:")
//   fmt.Println(result)
//
// Using the Client with Context
//
// The service's client also provides methods to make API requests with a Context
// value. This allows you to control the timeout, and cancellation of pending
// requests. These methods also take request Option as variadic parameter to apply
// additional configuration to the API request.
//
//   ctx := context.Background()
//
//   result, err := svc.PostContentWithContext(ctx, params)
//
// See the request package documentation for more information on using Context pattern
// with the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/request/
package lexruntimeservice
