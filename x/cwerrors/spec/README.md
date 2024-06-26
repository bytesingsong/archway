# CWErrors

This module is the error handling mechanism for the protocol to let the contract know of errors encountered

## Concepts

[callback](../../callback/spec/README.md) module and [cwica](../../cwica/spec/README.md) module both rely on protocol invoking contract execution. In case, there are any errors generated by the protocol, there needs to be a way to let the contract know of this error. The module provides a standard way of error catching from the contract side.

The module provides two diffferent ways that the error is communicated.

### 1. Errors saved in state (default)

Whenever a contract relevant error is encountered by the protocol, the module stores the error and associates with the contract address. This is stored in the chain state for `x` number of blocks (The `x` value is a module parameter. [See more](./01_state.md)) after which the error is automatically pruned. These stored errors are queryable by the contract using stargate queries.

### 2. Errors sudo callback

If a contract is subscribed to errors, then the contract will be invoked at the sudo entrypoint like so.

```rust
#[cw_serde]
pub enum SudoMsg  {
    Error {
        module_name: String, // The name of the module which generated the error
        error_code: u32, // module specific error code
        contract_address: String, // the contract address which is associated with the error; the contract receiving the callback
        input_payload: String, // any relevant input payload which caused the error
        error_message: String, // the relevant error message
    }
}
```

The subscriptions are an opt-in feature where the contractadmin/owner has to subscribe to the feature by paying the relevant fees. [See more](01_state.md) The subscription is valid for `x` number of blocks where `x` is decided by the module param. The subscription cannot be cancelled but can be extended by attempting to subscribe again.

When an error is received for a contract with the subscripiton, the module stores the errors in its transient store and executes the Sudo calls at the end block, by reading from the transient store.

## How to use in another module

If any module would like to integrate the x/cwerros functionality, add the following snippet into the `x/<yourmodule>/types/expected_keepers.go` 

```go
// ErrorsKeeper defines the expected interface needed to interact with the cwerrors module.
type ErrorsKeeper interface {
	// SetError records a sudo error for a contract
	SetError(ctx sdk.Context, sudoErr cwerrortypes.SudoError) error
}
```

This keeper would need to be passed in from `app.go` and stored by the module during init. When the module encounters an error it would like to be reported to the x/cwerrors module, it can execute the following snippet.

```go
sudoerr := cwerrortypes.SudoError {
		ModuleName:      ModuleName,
		ErrorCode:       int32(errorCode),
		ContractAddress: contractAddr,
		InputPayload:    inputPayload,
		ErrorMessage:    errMsg,
}
err := yourmodulekeeper.errorsKeeper.SetError(ctx, sudoerr)
```

## Contents

1. [State](./01_state.md)
2. [Messages](./02_messages.md)
3. [End Block](./03_end_block.md)
4. [Events](./04_events.md)
5. [Client](./05_client.md)
6. [Module Errors](./06_errors.md)

## References

1. [RFC: x/cwerrors module](https://github.com/orgs/archway-network/discussions/35)
2. [AIP: x/cwerrors module](https://github.com/archway-network/archway/issues/544)
