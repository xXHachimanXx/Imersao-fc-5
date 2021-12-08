import { Table, Model } from "sequelize-typescript";

@Table({
    tableName: 'account',
    createdAt: 'created_at',
    updatedAt: 'updated_at',
})
export class Account extends Model {
    
}
