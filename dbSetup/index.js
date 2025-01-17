const pg = require("pg")
const fs = require("fs")
const path = require("path")
const { Client } = pg
const client = new Client({
    connectionString: "postgresql://postgres:parsa@12345@localhost:5432/beautychi_dev"
})

async function f() {

    await client.connect()
    const schemaPath = path.join(__dirname, "schema.sql");

    try {
        const dropQuery = await client.query(`drop table if exists public.product_specifications cascade;

    drop table if exists public.product_image cascade;

    drop table if exists public.product cascade;

    drop table if exists public.sub_category cascade;

    drop table if exists public.category cascade;

    drop table if exists public.brand cascade;

    drop table if exists public.customer cascade;

    drop table if exists public.customer_address cascade;
    drop table if exists public.product_review cascade;

    `)

        const schema = fs.readFileSync(schemaPath, "utf8");
        const res = await client.query(schema)
        console.log(`Tables Droped And Recreated  Successfully!`)
    } catch (err) {
        console.error(err);
        process.exit(1);
    } finally {
        await client.end()
    }
}f()