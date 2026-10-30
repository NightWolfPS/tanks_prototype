components {
  id: "tank_spawner"
  component: "/main/scripts/tank_spawner.script"
}
components {
  id: "death"
  component: "/main/sounds/death.sound"
}
embedded_components {
  id: "tank_collection_factory"
  type: "collectionfactory"
  data: "prototype: \"/main/objects/tank.collection\"\n"
  "load_dynamically: true\n"
  "dynamic_prototype: true\n"
  ""
}
