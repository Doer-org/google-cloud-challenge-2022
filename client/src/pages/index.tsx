import { useState } from 'react';
import { MyHead } from '../components/templates/shared/Head/MyHead';
import { BasicTemplate } from '../components/templates/shared/BasicTemplate';
import { TypoWrapper } from '../components/atoms/text/TypoWrapper';
import { AuthLinks } from '../components/molecules/AuthLinks';
import style from '../styles/title.module.css';
export default function Home() {
  const [auth, setAuth] = useState(false);
  console.log("HELLOOOOOOOOO",process.env.NEXT_PUBLIC_GOOGLE_MAP_API,"???")
  return (
    <>
      <MyHead title="すきーま" description="すきーまの説明" />
      <BasicTemplate className="text-center">
        <TypoWrapper size="large" line="bold">
          <h1 className={style.title}>すきーま</h1>
        </TypoWrapper>
        <AuthLinks auth={auth} changeAuth={setAuth} />
      </BasicTemplate>
    </>
  );
}
