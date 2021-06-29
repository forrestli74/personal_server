declare module "*.ejs" {
	const value: (locals: any) => string;
	export = value;
}
declare module "*.proto" {
	const value: string;
	export = value;
}