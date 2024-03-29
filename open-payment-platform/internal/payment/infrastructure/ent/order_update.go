// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cloudwego/biz-demo/open-payment-platform/internal/payment/infrastructure/ent/order"
	"github.com/cloudwego/biz-demo/open-payment-platform/internal/payment/infrastructure/ent/predicate"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetMerchantID sets the "merchant_id" field.
func (ou *OrderUpdate) SetMerchantID(s string) *OrderUpdate {
	ou.mutation.SetMerchantID(s)
	return ou
}

// SetChannel sets the "channel" field.
func (ou *OrderUpdate) SetChannel(s string) *OrderUpdate {
	ou.mutation.SetChannel(s)
	return ou
}

// SetPayWay sets the "pay_way" field.
func (ou *OrderUpdate) SetPayWay(s string) *OrderUpdate {
	ou.mutation.SetPayWay(s)
	return ou
}

// SetOutOrderNo sets the "out_order_no" field.
func (ou *OrderUpdate) SetOutOrderNo(s string) *OrderUpdate {
	ou.mutation.SetOutOrderNo(s)
	return ou
}

// SetTotalAmount sets the "total_amount" field.
func (ou *OrderUpdate) SetTotalAmount(u uint64) *OrderUpdate {
	ou.mutation.ResetTotalAmount()
	ou.mutation.SetTotalAmount(u)
	return ou
}

// AddTotalAmount adds u to the "total_amount" field.
func (ou *OrderUpdate) AddTotalAmount(u int64) *OrderUpdate {
	ou.mutation.AddTotalAmount(u)
	return ou
}

// SetBody sets the "body" field.
func (ou *OrderUpdate) SetBody(s string) *OrderUpdate {
	ou.mutation.SetBody(s)
	return ou
}

// SetOrderStatus sets the "order_status" field.
func (ou *OrderUpdate) SetOrderStatus(i int8) *OrderUpdate {
	ou.mutation.ResetOrderStatus()
	ou.mutation.SetOrderStatus(i)
	return ou
}

// AddOrderStatus adds i to the "order_status" field.
func (ou *OrderUpdate) AddOrderStatus(i int8) *OrderUpdate {
	ou.mutation.AddOrderStatus(i)
	return ou
}

// SetAuthCode sets the "auth_code" field.
func (ou *OrderUpdate) SetAuthCode(s string) *OrderUpdate {
	ou.mutation.SetAuthCode(s)
	return ou
}

// SetWxAppid sets the "wx_appid" field.
func (ou *OrderUpdate) SetWxAppid(s string) *OrderUpdate {
	ou.mutation.SetWxAppid(s)
	return ou
}

// SetSubOpenid sets the "sub_openid" field.
func (ou *OrderUpdate) SetSubOpenid(s string) *OrderUpdate {
	ou.mutation.SetSubOpenid(s)
	return ou
}

// SetJumpURL sets the "jump_url" field.
func (ou *OrderUpdate) SetJumpURL(s string) *OrderUpdate {
	ou.mutation.SetJumpURL(s)
	return ou
}

// SetNotifyURL sets the "notify_url" field.
func (ou *OrderUpdate) SetNotifyURL(s string) *OrderUpdate {
	ou.mutation.SetNotifyURL(s)
	return ou
}

// SetClientIP sets the "client_ip" field.
func (ou *OrderUpdate) SetClientIP(s string) *OrderUpdate {
	ou.mutation.SetClientIP(s)
	return ou
}

// SetAttach sets the "attach" field.
func (ou *OrderUpdate) SetAttach(s string) *OrderUpdate {
	ou.mutation.SetAttach(s)
	return ou
}

// SetOrderExpiration sets the "order_expiration" field.
func (ou *OrderUpdate) SetOrderExpiration(s string) *OrderUpdate {
	ou.mutation.SetOrderExpiration(s)
	return ou
}

// SetExtendParams sets the "extend_params" field.
func (ou *OrderUpdate) SetExtendParams(s string) *OrderUpdate {
	ou.mutation.SetExtendParams(s)
	return ou
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ou.hooks) == 0 {
		affected, err = ou.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ou.mutation = mutation
			affected, err = ou.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ou.hooks) - 1; i >= 0; i-- {
			if ou.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ou.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ou.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: order.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.MerchantID(); ok {
		_spec.SetField(order.FieldMerchantID, field.TypeString, value)
	}
	if value, ok := ou.mutation.Channel(); ok {
		_spec.SetField(order.FieldChannel, field.TypeString, value)
	}
	if value, ok := ou.mutation.PayWay(); ok {
		_spec.SetField(order.FieldPayWay, field.TypeString, value)
	}
	if value, ok := ou.mutation.OutOrderNo(); ok {
		_spec.SetField(order.FieldOutOrderNo, field.TypeString, value)
	}
	if value, ok := ou.mutation.TotalAmount(); ok {
		_spec.SetField(order.FieldTotalAmount, field.TypeUint64, value)
	}
	if value, ok := ou.mutation.AddedTotalAmount(); ok {
		_spec.AddField(order.FieldTotalAmount, field.TypeUint64, value)
	}
	if value, ok := ou.mutation.Body(); ok {
		_spec.SetField(order.FieldBody, field.TypeString, value)
	}
	if value, ok := ou.mutation.OrderStatus(); ok {
		_spec.SetField(order.FieldOrderStatus, field.TypeInt8, value)
	}
	if value, ok := ou.mutation.AddedOrderStatus(); ok {
		_spec.AddField(order.FieldOrderStatus, field.TypeInt8, value)
	}
	if value, ok := ou.mutation.AuthCode(); ok {
		_spec.SetField(order.FieldAuthCode, field.TypeString, value)
	}
	if value, ok := ou.mutation.WxAppid(); ok {
		_spec.SetField(order.FieldWxAppid, field.TypeString, value)
	}
	if value, ok := ou.mutation.SubOpenid(); ok {
		_spec.SetField(order.FieldSubOpenid, field.TypeString, value)
	}
	if value, ok := ou.mutation.JumpURL(); ok {
		_spec.SetField(order.FieldJumpURL, field.TypeString, value)
	}
	if value, ok := ou.mutation.NotifyURL(); ok {
		_spec.SetField(order.FieldNotifyURL, field.TypeString, value)
	}
	if value, ok := ou.mutation.ClientIP(); ok {
		_spec.SetField(order.FieldClientIP, field.TypeString, value)
	}
	if value, ok := ou.mutation.Attach(); ok {
		_spec.SetField(order.FieldAttach, field.TypeString, value)
	}
	if value, ok := ou.mutation.OrderExpiration(); ok {
		_spec.SetField(order.FieldOrderExpiration, field.TypeString, value)
	}
	if value, ok := ou.mutation.ExtendParams(); ok {
		_spec.SetField(order.FieldExtendParams, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderMutation
}

// SetMerchantID sets the "merchant_id" field.
func (ouo *OrderUpdateOne) SetMerchantID(s string) *OrderUpdateOne {
	ouo.mutation.SetMerchantID(s)
	return ouo
}

// SetChannel sets the "channel" field.
func (ouo *OrderUpdateOne) SetChannel(s string) *OrderUpdateOne {
	ouo.mutation.SetChannel(s)
	return ouo
}

// SetPayWay sets the "pay_way" field.
func (ouo *OrderUpdateOne) SetPayWay(s string) *OrderUpdateOne {
	ouo.mutation.SetPayWay(s)
	return ouo
}

// SetOutOrderNo sets the "out_order_no" field.
func (ouo *OrderUpdateOne) SetOutOrderNo(s string) *OrderUpdateOne {
	ouo.mutation.SetOutOrderNo(s)
	return ouo
}

// SetTotalAmount sets the "total_amount" field.
func (ouo *OrderUpdateOne) SetTotalAmount(u uint64) *OrderUpdateOne {
	ouo.mutation.ResetTotalAmount()
	ouo.mutation.SetTotalAmount(u)
	return ouo
}

// AddTotalAmount adds u to the "total_amount" field.
func (ouo *OrderUpdateOne) AddTotalAmount(u int64) *OrderUpdateOne {
	ouo.mutation.AddTotalAmount(u)
	return ouo
}

// SetBody sets the "body" field.
func (ouo *OrderUpdateOne) SetBody(s string) *OrderUpdateOne {
	ouo.mutation.SetBody(s)
	return ouo
}

// SetOrderStatus sets the "order_status" field.
func (ouo *OrderUpdateOne) SetOrderStatus(i int8) *OrderUpdateOne {
	ouo.mutation.ResetOrderStatus()
	ouo.mutation.SetOrderStatus(i)
	return ouo
}

// AddOrderStatus adds i to the "order_status" field.
func (ouo *OrderUpdateOne) AddOrderStatus(i int8) *OrderUpdateOne {
	ouo.mutation.AddOrderStatus(i)
	return ouo
}

// SetAuthCode sets the "auth_code" field.
func (ouo *OrderUpdateOne) SetAuthCode(s string) *OrderUpdateOne {
	ouo.mutation.SetAuthCode(s)
	return ouo
}

// SetWxAppid sets the "wx_appid" field.
func (ouo *OrderUpdateOne) SetWxAppid(s string) *OrderUpdateOne {
	ouo.mutation.SetWxAppid(s)
	return ouo
}

// SetSubOpenid sets the "sub_openid" field.
func (ouo *OrderUpdateOne) SetSubOpenid(s string) *OrderUpdateOne {
	ouo.mutation.SetSubOpenid(s)
	return ouo
}

// SetJumpURL sets the "jump_url" field.
func (ouo *OrderUpdateOne) SetJumpURL(s string) *OrderUpdateOne {
	ouo.mutation.SetJumpURL(s)
	return ouo
}

// SetNotifyURL sets the "notify_url" field.
func (ouo *OrderUpdateOne) SetNotifyURL(s string) *OrderUpdateOne {
	ouo.mutation.SetNotifyURL(s)
	return ouo
}

// SetClientIP sets the "client_ip" field.
func (ouo *OrderUpdateOne) SetClientIP(s string) *OrderUpdateOne {
	ouo.mutation.SetClientIP(s)
	return ouo
}

// SetAttach sets the "attach" field.
func (ouo *OrderUpdateOne) SetAttach(s string) *OrderUpdateOne {
	ouo.mutation.SetAttach(s)
	return ouo
}

// SetOrderExpiration sets the "order_expiration" field.
func (ouo *OrderUpdateOne) SetOrderExpiration(s string) *OrderUpdateOne {
	ouo.mutation.SetOrderExpiration(s)
	return ouo
}

// SetExtendParams sets the "extend_params" field.
func (ouo *OrderUpdateOne) SetExtendParams(s string) *OrderUpdateOne {
	ouo.mutation.SetExtendParams(s)
	return ouo
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	var (
		err  error
		node *Order
	)
	if len(ouo.hooks) == 0 {
		node, err = ouo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ouo.mutation = mutation
			node, err = ouo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ouo.hooks) - 1; i >= 0; i-- {
			if ouo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ouo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ouo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Order)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from OrderMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   order.Table,
			Columns: order.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: order.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Order.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.MerchantID(); ok {
		_spec.SetField(order.FieldMerchantID, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Channel(); ok {
		_spec.SetField(order.FieldChannel, field.TypeString, value)
	}
	if value, ok := ouo.mutation.PayWay(); ok {
		_spec.SetField(order.FieldPayWay, field.TypeString, value)
	}
	if value, ok := ouo.mutation.OutOrderNo(); ok {
		_spec.SetField(order.FieldOutOrderNo, field.TypeString, value)
	}
	if value, ok := ouo.mutation.TotalAmount(); ok {
		_spec.SetField(order.FieldTotalAmount, field.TypeUint64, value)
	}
	if value, ok := ouo.mutation.AddedTotalAmount(); ok {
		_spec.AddField(order.FieldTotalAmount, field.TypeUint64, value)
	}
	if value, ok := ouo.mutation.Body(); ok {
		_spec.SetField(order.FieldBody, field.TypeString, value)
	}
	if value, ok := ouo.mutation.OrderStatus(); ok {
		_spec.SetField(order.FieldOrderStatus, field.TypeInt8, value)
	}
	if value, ok := ouo.mutation.AddedOrderStatus(); ok {
		_spec.AddField(order.FieldOrderStatus, field.TypeInt8, value)
	}
	if value, ok := ouo.mutation.AuthCode(); ok {
		_spec.SetField(order.FieldAuthCode, field.TypeString, value)
	}
	if value, ok := ouo.mutation.WxAppid(); ok {
		_spec.SetField(order.FieldWxAppid, field.TypeString, value)
	}
	if value, ok := ouo.mutation.SubOpenid(); ok {
		_spec.SetField(order.FieldSubOpenid, field.TypeString, value)
	}
	if value, ok := ouo.mutation.JumpURL(); ok {
		_spec.SetField(order.FieldJumpURL, field.TypeString, value)
	}
	if value, ok := ouo.mutation.NotifyURL(); ok {
		_spec.SetField(order.FieldNotifyURL, field.TypeString, value)
	}
	if value, ok := ouo.mutation.ClientIP(); ok {
		_spec.SetField(order.FieldClientIP, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Attach(); ok {
		_spec.SetField(order.FieldAttach, field.TypeString, value)
	}
	if value, ok := ouo.mutation.OrderExpiration(); ok {
		_spec.SetField(order.FieldOrderExpiration, field.TypeString, value)
	}
	if value, ok := ouo.mutation.ExtendParams(); ok {
		_spec.SetField(order.FieldExtendParams, field.TypeString, value)
	}
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
