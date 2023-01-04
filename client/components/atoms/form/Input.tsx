type TProps<T> = {
  type: React.HTMLInputTypeAttribute;
  label: string;
  content: string | number;
  changeContent: (content: T) => void;
};
export const Input = <T,>({
  type,
  label,
  content,
  changeContent,
}: TProps<T>) => {
  return (
    <div className="grid grid-cols-1 my-5">
      <label htmlFor={label}>{label}</label>
      <input
        type={type}
        className=" text-slate-900 py-1 px-2 rounded-sm"
        id={label}
        onChange={(e) => changeContent(e.target.value as T)}
        value={content}
        required={true}
      />
    </div>
  );
};
