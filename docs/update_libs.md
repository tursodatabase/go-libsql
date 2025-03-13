Install `cross` on your machine:

```
cargo install cross
```

Then build `libsql_experimental` library:

```
git clone https://github.com/tursodatabase/libsql.git
cd bindings/c
cross build --release --target x86_64-unknown-linux-gnu
cross build --release --target aarch64-unknown-linux-gnu
```

Finally copy library to `libs` directory of this repo:

```
libsql/target/<architecture>/release/libsql_experimental.a
```
