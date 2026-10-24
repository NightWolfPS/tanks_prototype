components {
  id: "projectile"
  component: "/main/scripts/projectile.script"
  properties {
    id: "speed"
    value: "750.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "lifetime"
    value: "1.5"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "damage"
    value: "50.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "min_damage"
    value: "1.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
components {
  id: "impact_proj"
  component: "/main/sounds/impact_proj.sound"
}
components {
  id: "impact"
  component: "/main/sounds/impact.sound"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"anim\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/assets/tank/projectile.tilesource\"\n"
  "}\n"
  ""
  position {
    z: 0.1
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"projectile\"\n"
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
  "  }\n"
  "  data: 2.0\n"
  "  data: 7.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "impact_factory"
  type: "factory"
  data: "prototype: \"/main/objects/impact.go\"\n"
  ""
}
