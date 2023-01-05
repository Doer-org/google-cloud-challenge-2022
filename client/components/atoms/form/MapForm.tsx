import { GoogleMap, LoadScript, Marker } from '@react-google-maps/api';
import { useState } from 'react';
export const MapForm = () => {
  const [pos, setPos] = useState<any>();
  const containerStyle = {
    width: '100%',
    height: '300px',
  };

  const center = {
    lat: 35.69575,
    lng: 139.77521,
  };
  console.log(pos);
  return (
    <LoadScript googleMapsApiKey={`${process.env.NEXT_PUBLIC_GOOGLE_MAP_API}`}>
      <GoogleMap
        mapContainerStyle={containerStyle}
        center={center}
        zoom={17}
        onClick={(e) => {
          setPos({
            lat: Number(e.latLng?.lat()),
            lng: Number(e.latLng?.lng()),
          });
        }}
      >
        <Marker position={pos} />
      </GoogleMap>
    </LoadScript>
  );
};
