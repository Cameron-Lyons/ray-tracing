# Ray Tracing

A ray tracer written in Go, based on [Ray Tracing in One Weekend](https://raytracing.github.io/) and [Ray Tracing: The Next Week](https://raytracing.github.io/books/RayTracingTheNextWeek.html).

## Features

- **Materials**: Lambertian (diffuse), metal (reflective with configurable fuzz), dielectric (glass with refraction), and diffuse light (emissive)
- **Textures**: Solid colors, checkerboard patterns, Perlin noise, and PNG image textures
- **Geometry**: Spheres, moving spheres, axis-aligned rectangles, and boxes
- **Camera**: Configurable field of view, depth of field, and motion blur
- **Volumetrics**: Constant density media for smoke and fog effects
- **Acceleration**: Bounding Volume Hierarchy (BVH) for fast ray-object intersection
- **Transformations**: Translation and Y-axis rotation
- **Rendering**: Multi-threaded scanline rendering with gamma correction

## Scenes

Two built-in scenes are available, selectable via the `scene` variable in `display_image.go`:

1. **Random scene** - A field of randomly placed spheres with various materials on a checkerboard ground, using Perlin noise and metal textures
2. **Cornell smoke** - A Cornell box with volumetric smoke and fog boxes

## Building and Running

```
go build -o raytracer .
./raytracer
```

Output is written to `output.png`.
