import { MyHead } from '../components/templates/shared/Head/MyHead';
import { BasicTemplate } from '../components/templates/shared/BasicTemplate';
import { TypoWrapper } from '../components/atoms/text/TypoWrapper';
import style from '../styles/title.module.css';
import { Button } from '../components/atoms/text/Button';
import { useUserInfoStore } from '../store/userStore';
import { createUser } from '../core/api/user/create';
import Link from 'next/link';
export default function Home() {
  const { userInfo, setUserInfo } = useUserInfoStore();
  const signIn = createUser(
    (ok) => {
      setUserInfo({
        userId: ok.id,
        userName: ok.name,
        icon: ok.icon ?? '',
      });
    },
    (err) => {}
  );
  return (
    <>
      <MyHead title="すきーま" description="すきーまの説明" />
      <BasicTemplate className="text-center">
        <TypoWrapper size="large" line="bold">
          <h1 className={style.title}>すきーま</h1>
        </TypoWrapper>
        <div className="my-10 grid grid-cols-1 gap-2">
          <Link
            href={`${process.env.NEXT_PUBLIC_SERVER_URL}/auth/login?redirect_url=http://localhost:3000`}
          >
            ログイン
          </Link>
        </div>
      </BasicTemplate>
    </>
  );
}
