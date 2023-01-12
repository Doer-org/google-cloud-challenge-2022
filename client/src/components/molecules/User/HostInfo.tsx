import React, { useEffect, useState } from 'react';
import { getUserInfo } from '../../../core/api/user/getInfo';
type TProps = {
  userId: string;
};
export const HostInfo = ({ userId }: TProps) => {
  const getUser = getUserInfo(
    (response) => {},
    (error) => {}
  );
  const [user, setUser] = useState();
  useEffect(() => {
    const userResponse = getUser(userId);
  }, []);
  return <div>HostInfo</div>;
};
