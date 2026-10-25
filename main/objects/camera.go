components {
  id: "camera_script"
  component: "/main/scripts/camera_script.script"
}
embedded_components {
  id: "camera"
  type: "camera"
  data: "aspect_ratio: 1.0\n"
  "fov: 0.7854\n"
  "near_z: 0.0\n"
  "far_z: 1000.0\n"
  "orthographic_projection: 1\n"
  "orthographic_zoom: 0.8\n"
  "orthographic_mode: ORTHO_MODE_AUTO_COVER\n"
  ""
}
