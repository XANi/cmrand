/*

cmrand provides convenient wrapper for generating math.rand Source() that uses crypto/rand to get its numbers

Usage:

    // to get *math.Rand with crypto/rand default source
    rnd := cmrand.New()
    // to just get source
    rndSrc := cmrand.NewSource(0) // seed is ignored

*/


package cmrand
