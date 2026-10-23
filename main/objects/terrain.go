components {
  id: "terrain1"
  component: "/main/objects/arena.tilemap"
}
components {
  id: "collisions"
  component: "/main/objects/collisions.tilemap"
  position {
    z: 0.3
  }
}
components {
  id: "houses"
  component: "/main/assets/houses/houses.tilemap"
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"/main/objects/collisions.tilemap\"\n"
  "type: COLLISION_OBJECT_TYPE_STATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"default\"\n"
  "mask: \"tank\"\n"
  ""
}
