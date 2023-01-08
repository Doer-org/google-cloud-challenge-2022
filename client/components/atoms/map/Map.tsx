import { GoogleMap, LoadScript, MarkerF } from '@react-google-maps/api';
import { ContainerStyle, Options, type TMapPosition } from './MapBasicInfo';
export const Map = (mapPosition: TMapPosition | null) => {
  return (
    <GoogleMap
      mapContainerStyle={ContainerStyle}
      center={mapPosition !== null ? mapPosition : undefined}
      zoom={15}
      options={Options}
    >
      {mapPosition ? <MarkerF position={mapPosition} /> : <></>}
    </GoogleMap>
  );
};
