import { GoogleMap, MarkerF } from '@react-google-maps/api';
import { useEffect, useState } from 'react';
import {
  ContainerStyle,
  Options,
  type TMapPosition,
} from '../map/MapBasicInfo';
export const MapForm = () => {
  const [pos, setPos] = useState<TMapPosition | null>(null);
  const [current, setCurrent] = useState<TMapPosition>({
    lat: 35.6809591,
    lng: 139.7673068,
  });
  useEffect(() => {
    navigator.geolocation.getCurrentPosition(
      (position) => {
        setCurrent({
          lat: position.coords.latitude,
          lng: position.coords.longitude,
        });
      },
      (err) => {
        throw new Error(err.message);
      }
    );
  }, []);

  return (
    <GoogleMap
      mapContainerStyle={ContainerStyle}
      center={current !== null ? current : undefined}
      zoom={15}
      onClick={(e) => {
        setPos({
          lat: Number(e.latLng?.lat()),
          lng: Number(e.latLng?.lng()),
        });
      }}
      options={Options}
    >
      {pos ? <MarkerF position={pos} /> : <></>}
    </GoogleMap>
  );
};
