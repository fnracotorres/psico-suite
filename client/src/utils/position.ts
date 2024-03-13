export enum Kind {
  Greater = "greater",
  Smaller = "smaller",
}

export enum Coordinate {
  X = "posX",
  Y = "posY",
}

export function getClosestToConstant<T extends Record<Coordinate, number>>(
  position: T,
  array: T[],
  kind: Kind,
  coordinate: Coordinate,
) {
  let closest: T | null = null;
  let minDistance = Infinity;
  const constant = position[coordinate];
  console.log(constant);

  for (const point of array) {
    const distance = Math.abs(point[coordinate] - constant);
    if (
      (kind === Kind.Greater &&
        typeof point[coordinate] === "number" &&
        point[coordinate] >= constant &&
        distance < minDistance) ||
      (kind === Kind.Smaller &&
        typeof point[coordinate] === "number" &&
        point[coordinate] <= constant &&
        distance < minDistance)
    ) {
      closest = point;
      minDistance = distance;
    }
  }

  console.log(closest);

  return closest;
}
