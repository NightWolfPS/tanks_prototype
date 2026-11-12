components {
  id: "shoot"
  component: "/main/sounds/shoot.sound"
}
components {
  id: "heavy_turret_controller"
  component: "/main/scripts/heavy_turret_controller.script"
}
components {
  id: "turret1"
  component: "/main/sounds/turret.sound"
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
  data: "prototype: \"/main/objects/projectiles/tank_turret_projectile.go\"\n"
  ""
}
embedded_components {
  id: "turret"
  type: "sprite"
  data: "default_animation: \"anim\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 169.0\n"
  "  y: 520.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/tank/heavy_turret.tilesource\"\n"
  "}\n"
  ""
  position {
    y: 11.0
    z: 0.5
  }
  scale {
    x: 0.2
    y: 0.2
  }
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
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/sprites1/pivot.tilesource\"\n"
  "}\n"
  ""
  position {
    y: 54.0
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
  "      y: 38.0\n"
  "      z: 0.1\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 5.0\n"
  "  data: 20.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
