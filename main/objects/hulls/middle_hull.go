components {
  id: "death"
  component: "/main/sounds/death.sound"
}
components {
  id: "engine"
  component: "/main/sounds/engine.sound"
}
components {
  id: "engine_idle"
  component: "/main/sounds/engine_idle.sound"
}
components {
  id: "tank_controller"
  component: "/main/scripts/middle_hull_controller.script"
  properties {
    id: "turn_speed"
    value: "75.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"tank\"\n"
  "mask: \"world\"\n"
  "mask: \"tank\"\n"
  "mask: \"bot\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: -1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"collision\"\n"
  "  }\n"
  "  data: 30.0\n"
  "  data: 47.5\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "middle_hull"
  type: "sprite"
  data: "default_animation: \"anim\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 397.0\n"
  "  y: 522.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/tank/middle_hull.tilesource\"\n"
  "}\n"
  ""
  position {
    y: -1.0
    z: 0.1
  }
  scale {
    x: 0.2
    y: 0.2
  }
}
