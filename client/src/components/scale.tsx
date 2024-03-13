import { useKeyDown } from "../hooks/useKeyboard";
import { socket } from "../socket";
import { getClosestToConstant, Kind, Coordinate } from "../utils/position";
import { useEffect, useState } from "react";

interface Device {
  name: string;
  posX: number;
  posY: number;
}

interface Head {
  device: Device;
}

export function Scale() {
  const [devices, setDevices] = useState<null | Device[]>(null);

  useEffect(() => {
    function onScale(devices: null | Device[]) {
      setDevices(devices);
    }

    socket.on("scale", onScale);
    socket.emit("scale", {});

    return () => {
      socket.off("scale", onScale);
    };
  }, []);

  const [head, setHead] = useState<null | Head>(null);

  useEffect(() => {
    if (devices === null) return;
    if (devices.length === 0) return;

    setHead({ device: devices[devices.length - 1] });
  }, [devices]);

  useKeyDown(() => {
    if (devices === null) return;
    if (devices.length === 0) return;

    const device: null | Device = getClosestToConstant(
      head!.device,
      devices,
      Kind.Greater,
      Coordinate.Y,
    );

    if (device === null) return;

    setHead({ device });
  }, ["ArrowUp"]);

  useKeyDown(() => {
    if (devices === null) return;
    if (devices.length === 0) return;

    const device: null | Device = getClosestToConstant(
      head!.device,
      devices,
      Kind.Greater,
      Coordinate.X,
    );

    if (device === null) return;

    setHead({ device });

    console.log(device);
  }, ["ArrowRight"]);

  useKeyDown(() => {
    if (devices === null) return;
    if (devices.length === 0) return;

    const device: null | Device = getClosestToConstant(
      head!.device,
      devices,
      Kind.Smaller,
      Coordinate.Y,
    );

    if (device === null) return;

    setHead({ device });

    console.log(device);
  }, ["ArrowDown"]);

  useKeyDown(() => {
    if (devices === null) return;
    if (devices.length === 0) return;

    const device: null | Device = getClosestToConstant(
      head!.device,
      devices,
      Kind.Smaller,
      Coordinate.X,
    );

    if (device === null) return;

    setHead({ device });

    console.log(device);
  }, ["ArrowLeft"]);

  if (devices === null) return <div>Null</div>;

  return (
    <div>
      {devices.map((device, i) => (
        <Device key={i} device={device} />
      ))}
    </div>
  );
}

function Device({ device }: { device: Device }) {
  return <div>{device.name}</div>;
}
