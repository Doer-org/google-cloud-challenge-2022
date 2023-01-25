import { MyHead } from '../components/templates/shared/Head/MyHead';
import { BasicTemplate } from '../components/templates/shared/BasicTemplate';
import { TypoWrapper } from '../components/atoms/text/TypoWrapper';
import style from '../styles/title.module.css';
import { Button } from '../components/atoms/text/Button';
import { useUserInfoStore } from '../store/userStore';
import { createUser } from '../core/api/user/create';
export default function Home() {
  const { userInfo, setUserInfo } = useUserInfoStore();
  const signIn = createUser(
    (ok) => {
      setUserInfo({
        userId: ok.id,
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
          <Button
            className="m-1"
            onClick={() => {
              if (userInfo.userId === '') {
                signIn({
                  name: 'mahiro',
                  authenticated: true,
                });
              }
            }}
          >
            ログイン
          </Button>
          <Button
            className="m-1"
            onClick={() => {
              signIn({
                name: 'mahiro',
                authenticated: true,
              });
              console.log('create new account');
            }}
          >
            サインイン
          </Button>
        </div>
      </BasicTemplate>
    </>
  );
}
