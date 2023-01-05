import { useState } from 'react';
import { MyHead } from '../components/templates/shared/MyHead';
import { BasicTemplate } from '../components/templates/shared/BasicTemplate';
import { TypoWrapper } from '../components/atoms/TypoWrapper';
import { AuthLinks } from '../components/molecules/AuthLinks';

export default function Home() {
  const [auth, setAuth] = useState(false);
  return (
    <>
      <MyHead title="サービス名" description="サービスの説明" />
      <BasicTemplate className="text-center">
        <TypoWrapper size="large" line="bold">
          <h1>サービス名</h1>
        </TypoWrapper>
        <AuthLinks auth={auth} changeAuth={setAuth} />
      </BasicTemplate>
    </>
  );
}
