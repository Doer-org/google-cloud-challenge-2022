import { useEffect, useState } from 'react';
import { MapForm } from '../../components/atoms/form/MapForm';
import { TMapPosition } from '../../components/atoms/map/MapBasicInfo'; 
import { BasicTemplate } from '../../components/templates/shared/BasicTemplate'; 

export default function New() {  
  useEffect(() => { 
    console.log(window.location.origin)
  },[])
  const [location, setLocation] = useState<null | TMapPosition>(null);  
  return (
    <BasicTemplate className="text-center"> 
      <MapForm location={location} setLocation={setLocation} />
    </BasicTemplate>
  );
} 