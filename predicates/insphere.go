package predicates

import (
	"geGoMetry/r3"
	"math"
)

const isperrboundA = (16 + 224*epsilon) * epsilon

//InSphereFast : non robust implementation of the sphere method
func InSphereFast(pa, pb, pc, pd, pe r3.Vector) float64 {
	aex := pa.X - pe.X
	bex := pb.X - pe.X
	cex := pc.X - pe.X
	dex := pd.X - pe.X
	aey := pa.Y - pe.Y
	bey := pb.Y - pe.Y
	cey := pc.Y - pe.Y
	dey := pd.Y - pe.Y
	aez := pa.Z - pe.Z
	bez := pb.Z - pe.Z
	cez := pc.Z - pe.Z
	dez := pd.Z - pe.Z

	ab := aex*bey - bex*aey
	bc := bex*cey - cex*bey
	cd := cex*dey - dex*cey
	da := dex*aey - aex*dey
	ac := aex*cey - cex*aey
	bd := bex*dey - dex*bey

	abc := aez*bc - bez*ac + cez*ab
	bcd := bez*cd - cez*bd + dez*bc
	cda := cez*da + dez*ac + aez*cd
	dab := dez*ab + aez*bd + bez*da

	alift := aex*aex + aey*aey + aez*aez
	blift := bex*bex + bey*bey + bez*bez
	clift := cex*cex + cey*cey + cez*cez
	dlift := dex*dex + dey*dey + dez*dez

	return (clift*dab - dlift*abc) + (alift*bcd - blift*cda)
}

func inSphere(ax, ay, az, bx, by, bz, cx, cy, cz, dx, dy, dz, ex, ey, ez float64) float64 {
	aex := ax - ex
	bex := bx - ex
	cex := cx - ex
	dex := dx - ex
	aey := ay - ey
	bey := by - ey
	cey := cy - ey
	dey := dy - ey
	aez := az - ez
	bez := bz - ez
	cez := cz - ez
	dez := dz - ez

	aexbey := aex * bey
	bexaey := bex * aey
	ab := aexbey - bexaey
	bexcey := bex * cey
	cexbey := cex * bey
	bc := bexcey - cexbey
	cexdey := cex * dey
	dexcey := dex * cey
	cd := cexdey - dexcey
	dexaey := dex * aey
	aexdey := aex * dey
	da := dexaey - aexdey
	aexcey := aex * cey
	cexaey := cex * aey
	ac := aexcey - cexaey
	bexdey := bex * dey
	dexbey := dex * bey
	bd := bexdey - dexbey

	abc := aez*bc - bez*ac + cez*ab
	bcd := bez*cd - cez*bd + dez*bc
	cda := cez*da + dez*ac + aez*cd
	dab := dez*ab + aez*bd + bez*da

	alift := aex*aex + aey*aey + aez*aez
	blift := bex*bex + bey*bey + bez*bez
	clift := cex*cex + cey*cey + cez*cez
	dlift := dex*dex + dey*dey + dez*dez

	det := (clift*dab - dlift*abc) + (alift*bcd - blift*cda)

	aezplus := math.Abs(aez)
	bezplus := math.Abs(bez)
	cezplus := math.Abs(cez)
	dezplus := math.Abs(dez)
	aexbeyplus := math.Abs(aexbey)
	bexaeyplus := math.Abs(bexaey)
	bexceyplus := math.Abs(bexcey)
	cexbeyplus := math.Abs(cexbey)
	cexdeyplus := math.Abs(cexdey)
	dexceyplus := math.Abs(dexcey)
	dexaeyplus := math.Abs(dexaey)
	aexdeyplus := math.Abs(aexdey)
	aexceyplus := math.Abs(aexcey)
	cexaeyplus := math.Abs(cexaey)
	bexdeyplus := math.Abs(bexdey)
	dexbeyplus := math.Abs(dexbey)
	permanent := ((cexdeyplus+dexceyplus)*bezplus+(dexbeyplus+bexdeyplus)*cezplus+(bexceyplus+cexbeyplus)*dezplus)*alift +
		((dexaeyplus+aexdeyplus)*cezplus+(aexceyplus+cexaeyplus)*dezplus+(cexdeyplus+dexceyplus)*aezplus)*blift +
		((aexbeyplus+bexaeyplus)*dezplus+(bexdeyplus+dexbeyplus)*aezplus+(dexaeyplus+aexdeyplus)*bezplus)*clift +
		((bexceyplus+cexbeyplus)*aezplus+(cexaeyplus+aexceyplus)*bezplus+(aexbeyplus+bexaeyplus)*cezplus)*dlift

	errbound := isperrboundA * permanent
	if det > errbound || -det > errbound {
		return det
	}
	return -inSphereAdapt(ax, ay, az, bx, by, bz, cx, cy, cz, dx, dy, dz, ex, ey, ez, permanent)
}

func inSphereAdapt(ax, ay, az, bx, by, bz, cx, cy, cz, dx, dy, dz, ex, ey, ez, permanent float64) float64 {
	return 0.0
}
