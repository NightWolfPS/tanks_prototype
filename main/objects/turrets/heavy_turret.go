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
