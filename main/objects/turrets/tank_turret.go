components {
  id: "tank_turret_controller"
  component: "/main/scripts/tank_turret_controller.script"
  properties {
    id: "turret_turn_speed"
    value: "65.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "fire_cooldown"
    value: "4.5"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "muzzle_distance"
    value: "50.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
components {
  id: "shoot"
  component: "/main/sounds/shoot.sound"
}
components {
  id: "turret1"
  component: "/main/sounds/turret.sound"
}
embedded_components {
  id: "turret"
  type: "sprite"
  data: "default_animation: \"anim\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/tank/medium_turret.tilesource\"\n"
  "}\n"
  ""
  position {
    y: 24.0
  }
  scale {
    x: 0.2
    y: 0.2
  }
}
embedded_components {
  id: "projectile_factory"
  type: "factory"
  data: "prototype: \"/main/objects/projectiles/tank_turret_projectile.go\"\n"
  ""
}
embedded_components {
  id: "muzzle_flash_factory"
  type: "factory"
  data: "prototype: \"/main/objects/muzzle_flash.go\"\n"
  ""
}
