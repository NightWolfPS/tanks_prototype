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
    y: 23.0
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
  data: "prototype: \"/main/objects/projectile.go\"\n"
  ""
}
