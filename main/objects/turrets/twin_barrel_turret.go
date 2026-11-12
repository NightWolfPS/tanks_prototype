components {
  id: "shoot"
  component: "/main/sounds/shoot.sound"
}
components {
  id: "turret1"
  component: "/main/sounds/turret.sound"
}
components {
  id: "twin_barrel_turret_controller"
  component: "/main/scripts/twin_barrel_turret_controller.script"
  properties {
    id: "turret_turn_speed"
    value: "80.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "fire_cooldown"
    value: "0.75"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "muzzle_distance"
    value: "75.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "turret"
  type: "sprite"
  data: "default_animation: \"anim\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/tank/twin_barrel_turret.tilesource\"\n"
  "}\n"
  ""
  position {
    x: 2.0
    y: 24.0
    z: 0.5
  }
  scale {
    x: 0.2
    y: 0.2
  }
}
embedded_components {
  id: "muzzle_flash_factory"
  type: "factory"
  data: "prototype: \"/main/objects/muzzle_flash.go\"\n"
  ""
}
embedded_components {
  id: "projectile_factory"
  type: "factory"
  data: "prototype: \"/main/objects/projectiles/twin_barrel_turret_projectile.go\"\n"
  ""
}
embedded_components {
  id: "laser_factory"
  type: "factory"
  data: "prototype: \"/main/objects/laser.go\"\n"
  ""
}
embedded_components {
  id: "pivot"
  type: "sprite"
  data: "default_animation: \"anim\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 1.0\n"
  "  y: 1.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/sprites1/pivot.tilesource\"\n"
  "}\n"
  ""
  position {
    y: 28.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"turret\"\n"
  "mask: \"world\"\n"
  "mask: \"bot\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -7.0\n"
  "      y: 52.0\n"
  "      z: 0.1\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 8.0\n"
  "      y: 52.0\n"
  "      z: 0.1\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 3\n"
  "    count: 3\n"
  "  }\n"
  "  data: 4.0\n"
  "  data: 25.0\n"
  "  data: 10.0\n"
  "  data: 4.0\n"
  "  data: 25.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
