import { Button } from '../atoms/Button';
import { LinkTo } from '../atoms/LinkTo';
type TProps = {
  auth: boolean;
  changeAuth: (auth: boolean) => void;
};
export const AuthLinks = ({ auth, changeAuth }: TProps) => {
  return (
    <div className="my-10 grid grid-cols-1 gap-2">
      {!auth ? (
        <>
          <Button onClick={() => changeAuth(true)} className="m-1">
            ログイン
          </Button>
          <Button onClick={() => {}} className="m-1">
            サインイン
          </Button>
        </>
      ) : (
        <>
          <LinkTo href="/event/new" className="m-1">
            イベント募集ページへ
          </LinkTo>
          <Button onClick={() => changeAuth(false)} className="m-1">
            ログアウト
          </Button>
        </>
      )}
    </div>
  );
};
