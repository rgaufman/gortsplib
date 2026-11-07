# TetherX camera-compatibility integration

## Upstream sync, 7 November 2026

PROVEN: merged upstream main `e4283d1721bc46f55a177bed164d7b59a51b635d`.
The artificial GStreamer delay is now restricted to affected publisher versions
before1.16.3. A new table regression covers legacy/fixed/current versions and
missing, ambiguous and malformed headers.

OBSERVED: full native Linux ARM64 `go test -race -p 1 -count=1 -timeout 8m ./...`
passes on an isolated admitted CPU-only Spark job with Go1.27.2, including the
root networking suite. The sealed source archive SHA256 is
`a9fc3f0fed32153d4a3faa0d3f49d0398ddd90e196b5f8bba4b987db037e2ff7`.
Results were retrieved before the workstation reboot; they do not need repeating
for that reboot. No running camera service was changed.

This branch retains normalized session jitter (`df6c90d8e37f`) and integrates:

- [PR #1178](https://github.com/bluenviron/gortsplib/pull/1178), heads
  `13782df6502cb8f0210db5c00f1f12e260747633` and
  `03ccf9d7bee468862b1de4bad174a030e4cd979b`: use a single unambiguous FMTP
  attribute for a single-format description when its PT is mislabeled.
- [PR #1181](https://github.com/bluenviron/gortsplib/pull/1181), head
  `c99800a59b60095753970e9c21157a6aae4b036f`: a TCP SETUP response with an
  absent Transport header can accept the requested interleaved channels.

PROVEN from code/tests: ambiguous multi-format/FMTP descriptions remain strict;
the missing-header fallback is TCP-only. Malformed present headers and UDP still
go through the existing Transport parser. These are protocol compatibility
changes, not recorder clock or file policy.

OBSERVED local validation: race-enabled format/description packages, the new
`TestClientPlayNoTransportHeader` and existing session-jitter tests pass. The full
`pkg/...` and `internal/...` race suites pass too. This is not a claim of the
entire root multicast suite on macOS or deployment to a physical camera.

The matching MediaMTX integration records this library commit in its dependency
pin and documents the combined tests in `contrib/tetherx/validation`. Local
integration uses an explicit Go workspace; no machine-specific replacement path
is committed to the application's go.mod. Publish the library commit before
expecting a fresh remote checkout to resolve that new pin.
