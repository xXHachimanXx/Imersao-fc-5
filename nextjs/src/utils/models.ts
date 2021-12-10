export enum OrderStatus {
    Pending = "pending",
    Approved = "approved",
    Rejected = "rejected",
}

export const OrderStatusTranslate = {
    [OrderStatus.Pending]: 'Pending',
    [OrderStatus.Approved]: 'Approved',
    [OrderStatus.Rejected]: 'Rejected',
}