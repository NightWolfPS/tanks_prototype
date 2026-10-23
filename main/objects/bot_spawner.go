components {
  id: "bot_spawner"
  component: "/main/scripts/bot_spawner.script"
}
embedded_components {
  id: "bot_collection_factory"
  type: "collectionfactory"
  data: "prototype: \"/main/objects/bot.collection\"\n"
  ""
}
