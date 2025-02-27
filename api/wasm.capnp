using Go = import "/go.capnp";

@0x9419a7a54f76d35b;

$Go.package("wasm");
$Go.import("github.com/wetware/ww/internal/api/wasm");


interface Runtime extends(Executor(Context, Module)) {
    struct Context {
        src      @0 :Data;
        env      @1 :List(Field);
        
        stdin    @2 :IOStream.Provider;
        stdout   @3 :IOStream.Stream;
        stderr   @4 :IOStream.Stream;
        
        randSeed @5 :UInt64;

        struct Field {
            key   @0 :Text;
            value @1 :Text;
        }

        using IOStream = import "iostream.capnp";
    }

    interface Module extends(Waiter(Status)) {
        run   @0 () -> ();
        # Run the compiled WASM module in the present context.

        close @1 (exitCode :UInt32) -> ();
        # Close all the modules that have been initialized in this Runtime
        # with the provided exit code.  An error is returned if any module
        # returns an error when closed.

        struct Status {
            statusCode @0 :UInt32;
        }

        using Waiter = import "proc.capnp".Waiter;
    }
    
    using Executor = import "proc.capnp".Executor;
}
