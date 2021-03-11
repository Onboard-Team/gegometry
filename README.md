# geGoMetry:

geGoMetry is a converter from *step* to *glTF* written in goLang.
	- ***the work is still in progress***

### Installation & Setup: 
To install the dependencies, run:

```bash 
go get
```

The project is still in progress. for testing the current implemenations you could run 

```bash
cd `[whatever package you're interested in]`
go test
```


### Poject State: 
- [x] Creating the STEP reader.
- [x] Integrating the glTF exporter.
- [x] Supporting geometry primitives.
- [x] implementing 2D delaunay Triangulation.
- [x] implementing 3D delaunay Triangulation.
- [x] Integrating protobuf to integrate C++ libs
- [ ] Supporting NURBS.
- [ ] implementing Shewchuk's robust primitives.
- [ ] supporting the totality of the ISO 10303-21 templates.

###LICENCE:
MIT
