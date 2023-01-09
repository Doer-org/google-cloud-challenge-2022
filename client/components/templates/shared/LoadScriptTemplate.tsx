import { LoadScript } from '@react-google-maps/api';
import { useEffect, useState } from 'react';
type TProps = {
  children: JSX.Element;
};
export const LoadScriptTemplate = ({ children }: TProps) => {
  const [google, setGoogle] = useState<any>();
  useEffect(() => {
    setGoogle(window.google);
  }, []);
  return google === undefined ? (
    <LoadScript
      googleMapsApiKey={process.env.NEXT_PUBLIC_GOOGLE_MAP_API as string}
    >
      {children}
    </LoadScript>
  ) : (
    <>{children}</>
  );
};
