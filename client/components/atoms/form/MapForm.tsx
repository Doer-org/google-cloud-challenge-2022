import { GoogleMap, MarkerF, useJsApiLoader } from '@react-google-maps/api';
import { useEffect, useState } from 'react';
import {
  ContainerStyle,
  Options,
  type TMapPosition,
} from '../map/MapBasicInfo';
import { TypoWrapper } from '../text/TypoWrapper';
type TProps = {
  location: TMapPosition | null;
  setLocation: (position: TMapPosition) => void;
};
export const MapForm = ({ location, setLocation }: TProps) => {
  // TODO:atomsの責務を超えているのでmoleculesに分離するべき
  const [current, setCurrent] = useState<TMapPosition>({
    lat: 35.6809591,
    lng: 139.7673068,
  });
  useEffect(() => {
    navigator.geolocation.getCurrentPosition((position) => {
      setCurrent({
        lat: position.coords.latitude,
        lng: position.coords.longitude,
      });
    });
  }, []);
  const { isLoaded } = useJsApiLoader({
    id: 'google-map-script',
    googleMapsApiKey: process.env.NEXT_PUBLIC_GOOGLE_MAP_API as string,
  });
  return isLoaded && current ? (
    <GoogleMap
      mapContainerStyle={ContainerStyle}
      center={current !== null ? current : undefined}
      zoom={13}
      onClick={(e) => {
        setLocation({
          lat: Number(e.latLng?.lat()),
          lng: Number(e.latLng?.lng()),
        });
      }}
      options={Options}
    >
      {location ? <MarkerF position={location} /> : <></>}
    </GoogleMap>
  ) : (
    <TypoWrapper>
      <p>ローディング中</p>
    </TypoWrapper>
  );
};
