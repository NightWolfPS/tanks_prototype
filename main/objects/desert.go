components {
  id: "desert_map_dm"
  component: "/main/assets/desert_map_2/desert_map_dm.tilemap"
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_DYNAMIC\n"
  "mass: 1.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"world\"\n"
  "mask: \"tank\"\n"
  "mask: \"bot\"\n"
  ""
}
