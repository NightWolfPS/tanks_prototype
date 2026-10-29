components {
  id: "tank_controller"
  component: "/main/scripts/tank_controller.script"
  properties {
    id: "move_speed"
    value: "98.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "turn_speed"
    value: "45.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "max_hp"
    value: "100.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
components {
  id: "engine"
  component: "/main/sounds/engine.sound"
}
components {
  id: "death"
  component: "/main/sounds/death.sound"
}
components {
  id: "engine_idle"
  component: "/main/sounds/engine_idle.sound"
}
embedded_components {
  id: "tank_hull"
  type: "sprite"
  data: "default_animation: \"anim\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/tank/light_tank.tilesource\"\n"
  "}\n"
  ""
  position {
    z: 0.1
  }
  scale {
    x: 0.2
    y: 0.2
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
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "    id: \"collision\"\n"
  "  }\n"
  "  data: 36.5\n"
  "  data: 52.5\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
