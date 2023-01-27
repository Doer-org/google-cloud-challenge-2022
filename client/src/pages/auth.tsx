import { MyHead } from '../components/templates/shared/Head/MyHead';
import { BasicTemplate } from '../components/templates/shared/BasicTemplate';
import { TypoWrapper } from '../components/atoms/text/TypoWrapper';
import { Button } from '../components/atoms/text/Button';
import { useUserInfoStore } from '../store/userStore';
import { createUser } from '../core/api/user/create';
import Link from 'next/link';
import { TopPage } from '../components/molecules/Top/TopPage';
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
        <TopPage />
        <div className="my-10 grid grid-cols-1 gap-2">
          <Link
            href={`${process.env.NEXT_PUBLIC_SERVER_URL}/auth/login?redirect_url=${process.env.NEXT_PUBLIC_FRONT_URL}`}
          >
            ログイン
          </Link>
        </div>
      </BasicTemplate>
    </>
  );
}
