import { useEffect, useState } from 'react';
import { MapForm } from '../../components/atoms/form/MapForm';
import { TMapPosition } from '../../components/atoms/map/MapBasicInfo';
import { Button } from '../../components/atoms/text/Button';
import { BasicTemplate } from '../../components/templates/shared/BasicTemplate';
import { useNoticeStore } from '../../store/noticeStore';

export default function New() {
  useEffect(() => {
    console.log(window.location.origin);
  }, []);
  const { notice, changeNotice, resetNotice } = useNoticeStore();
  console.log(notice);
  const [location, setLocation] = useState<null | TMapPosition>(null);
  return (
    <BasicTemplate className="text-center">
      <Button
        onClick={() => {
          changeNotice({ type: 'Success', text: '成功やで' });
        }}
      >
        処理が成功ボタン
      </Button>
      <Button
        onClick={() => {
          changeNotice({ type: 'Error', text: '失敗やで' });
        }}
      >
        処理が失敗ボタン
      </Button>
      <Button onClick={() => resetNotice()}>リセットボタン</Button>
      <MapForm location={location} setLocation={setLocation} />
    </BasicTemplate>
  );
}
