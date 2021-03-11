package predicates

import "geGoMetry/r3"

func orient3dfast(a, b, c, d r3.Vector) float64 {
	adx := a.X - d.X
	bdx := b.X - d.X
	cdx := c.X - d.X
	ady := a.Y - d.Y
	bdy := b.Y - d.Y
	cdy := c.Y - d.Y
	adz := a.Z - d.Z
	bdz := b.Z - d.Z
	cdz := c.Z - d.Z

	return adx*(bdy*cdz-bdz*cdy) +
		bdx*(cdy*adz-cdz*ady) +
		cdx*(ady*bdz-adz*bdy)
}
