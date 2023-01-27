import { GoogleMap, MarkerF, useJsApiLoader } from '@react-google-maps/api';
import { TypoWrapper } from '../text/TypoWrapper';
import { ContainerStyle, Options, type TMapPosition } from './MapBasicInfo';
export const Map = (mapPosition: TMapPosition | null) => {
  const { isLoaded } = useJsApiLoader({
    id: 'google-map-script',
    googleMapsApiKey: process.env.NEXT_PUBLIC_GOOGLE_MAP_API as string,
  });
  return isLoaded ? (
    <div className="rounded-md">
      <GoogleMap
        mapContainerStyle={ContainerStyle}
        center={mapPosition !== null ? mapPosition : undefined}
        zoom={15}
        options={Options}
      >
        {mapPosition ? <MarkerF position={mapPosition} /> : <></>}
      </GoogleMap>
    </div>
  ) : (
    <TypoWrapper>
      <p>ローディング中</p>
    </TypoWrapper>
  );
};
