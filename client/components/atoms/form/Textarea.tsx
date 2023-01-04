type TProps = {
  label: string;
  content: string;
  changeContent: (content: string) => void;
};
export const Textarea = ({ label, content, changeContent }: TProps) => {
  return (
    <div className="grid grid-cols-1 my-5">
      <label htmlFor={label}>{label}</label>
      <textarea
        id={label}
        className="text-slate-900 py-1 px-2 rounded-sm max-h-32"
        onChange={(e) => changeContent(e.target.value)}
        value={content}
      />
    </div>
  );
};
