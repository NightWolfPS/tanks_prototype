components {
  id: "snowy_castle"
  component: "/main/assets/mapTiles/snowy_castle.tilemap"
}
components {
  id: "snowy_castle_collisions"
  component: "/main/assets/mapTiles/snowy_castle_collisions.tilemap"
  position {
    z: 0.5
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"/main/assets/mapTiles/snowy_castle_collisions.tilemap\"\n"
  "type: COLLISION_OBJECT_TYPE_STATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"default\"\n"
  "mask: \"tank\"\n"
  "mask: \"bot\"\n"
  "mask: \"turret\"\n"
  ""
}
