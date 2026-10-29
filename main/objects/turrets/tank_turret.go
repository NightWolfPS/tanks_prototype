components {
  id: "turret_controller"
  component: "/main/scripts/turret_controller.script"
  properties {
    id: "fire_cooldown"
    value: "6.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "muzzle_distance"
    value: "85.0"
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
  data: "prototype: \"/main/objects/projectile.go\"\n"
  ""
}
embedded_components {
  id: "muzzle_flash_factory"
  type: "factory"
  data: "prototype: \"/main/objects/muzzle_flash.go\"\n"
  ""
}
