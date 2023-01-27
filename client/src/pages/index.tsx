import { MyHead } from '../components/templates/shared/Head/MyHead';
import { BasicTemplate } from '../components/templates/shared/BasicTemplate';
import { LinkTo } from '../components/atoms/text/LinkTo';
import { useEffect } from 'react';
import { useUserInfoStore } from '../store/userStore';
import { TopPage } from '../components/molecules/Top/TopPage';
export default function Home() {
  const { setUserInfo } = useUserInfoStore();
  useEffect(() => {
    fetch(`${process.env.NEXT_PUBLIC_SERVER_URL}/auth/user`, {
      method: 'GET',
      credentials: 'include',
    }).then(async (ok) => {
      const body: {
        id: string;
        name: string;
        authenticated: boolean;
        mail: string;
        icon: string;
      } = await ok.json();
      // console.log(body)
      setUserInfo({
        userId: body.id,
        userName: body.name,
        icon: body.icon,
      });
    });
  }, []);
  return (
    <>
      <MyHead title="すきーま" description="すきーまの説明" />
      <BasicTemplate className="text-center">
        <TopPage />
        <div className="my-10 grid grid-cols-1 gap-5">
          <LinkTo href="/event/new" imgPath='/images/top/event.png' className="mx-12">
            イベント作成ページへ
          </LinkTo>
          <LinkTo href="/event" className="m-1">
            イベント一覧ページへ
          </LinkTo>
        </div>
      </BasicTemplate>
    </>
  );
}
