components {
  id: "garage_preview"
  component: "/main/scripts/garage_preview.script"
}
embedded_components {
  id: "light_hull_factory"
  type: "factory"
  data: "prototype: \"/main/objects/hulls/light_hull_prev.go\"\n"
  "load_dynamically: true\n"
  "dynamic_prototype: true\n"
  ""
}
embedded_components {
  id: "tank_turret_factory"
  type: "factory"
  data: "prototype: \"/main/objects/turrets/tank_turret_prev.go\"\n"
  "load_dynamically: true\n"
  "dynamic_prototype: true\n"
  ""
}
embedded_components {
  id: "twin_barrel_turret_factory"
  type: "factory"
  data: "prototype: \"/main/objects/turrets/twin_barrel_turret_prev.go\"\n"
  "load_dynamically: true\n"
  "dynamic_prototype: true\n"
  ""
}
embedded_components {
  id: "weight_hull_factory"
  type: "factory"
  data: "prototype: \"/main/objects/hulls/weight_hull_prev.go\"\n"
  "load_dynamically: true\n"
  "dynamic_prototype: true\n"
  ""
}
