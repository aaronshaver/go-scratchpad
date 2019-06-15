# derangement

## Description

I wrote this to practice Go and to solve a real life problem: I had framed artwork that I wanted to shuffle around my house such that every piece in the new arrangement was not in its starting location.

From https://en.wikipedia.org/wiki/Derangement:

*In combinatorial mathematics, a derangement is a permutation of the elements of a set, such that no element appears in its original position.*

## Example output

    C:\repos\go-scratchpad\derangement>go run main.go
    Type a string, then press <ENTER>. Press <ENTER> by itself to stop.
    -> PaintingA
    -> DrawingB
    -> WatercolorC
    -> FramedArtD
    -> TapestryE
    -> PieceOfPaperF
    -> ArtworkG
    ->
    [DrawingB WatercolorC TapestryE ArtworkG PieceOfPaperF FramedArtD PaintingA]
