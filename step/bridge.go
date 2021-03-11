package step

/* this file is a bridge between the step definitions and the
---converter definitions.
*/
/* Idea: the bridge could decide what and what not to execute
---making it possible to ignore data functions like cartesian
---and vertex points, axis2Placement3 and more..*/
func (parser *Parser) bridge(definition Definition) parseFn {
	switch definition.Type {
	case "CARTESIAN_POINT":
		/* skipping this definition for now */
		/* return parser.cartesian_point */
		return parser.skipUtility
	case "VERTEX_POINT":
		/* skipping this definition for now */
		/* return parser.vertexPoint */
		return parser.skipUtility
	case "VERTEX_LOOP":
		return parser.vertexLoop
	case "AXIS_2_PLACEMENT_3D":
		return parser.axis2Placement3D
	case "FACE":
		return parser.face
	case "B_SPLINE_SURFACE_WITH_KNOTS":
		return parser.bSplineSurfaceWithKnots
	case "ADVENCED_FACE":
		return parser.advancedFace
	case "CLOSED_SHELL":
		return parser.closedShell
	case "DIRECTION":
		return parser.direction
	case "PRODUCT_DEFINITION":
		return parser.productDefinition
	case "PRODUCT_DEFINITION_CONTEXT":
		return parser.productDefinitionContext
	case "PRODUCT_DEFINITION_FORMATION":
		return parser.productDefinitionFormation
	case "PRODUCT_DEFINITION_SHAPE":
		return parser.productDefinitionShape
	case "SPHERICAL_SURFACE":
		return parser.sphericalSurface
	case "TOROIDAL_SURFACE":
		return parser.toroidalSurface
	default:
		return parser.skipUtility
	}
}
