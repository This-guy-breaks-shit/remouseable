// This file is part of remouseable.
//
// remouseable is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License version 3 as published
// by the Free Software Foundation.
//
// remouseable is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with remouseable.  If not, see <https://www.gnu.org/licenses/>.

package remouseable

import (
	"fmt"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRuntimeEmptyStateMachine(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	d := NewMockDriver(ctrl)
	p := NewMockPositionScaler(ctrl)
	s := NewMockStateMachine(ctrl)
	rt := &Runtime{
		Driver:         d,
		PositionScaler: p,
		StateMachine:   s,
	}
	s.EXPECT().Next().Return(false).AnyTimes()
	s.EXPECT().Close().Return(nil)
	require.False(t, rt.Next())
	require.False(t, rt.Next())
	require.Nil(t, rt.Close())
}

func TestRuntimeStopsInErrorState(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	d := NewMockDriver(ctrl)
	p := NewMockPositionScaler(ctrl)
	s := NewMockStateMachine(ctrl)
	rt := &Runtime{
		Driver:         d,
		PositionScaler: p,
		StateMachine:   s,
		err:            fmt.Errorf("test"),
	}
	s.EXPECT().Close().Return(nil)
	require.False(t, rt.Next())
	require.False(t, rt.Next())
	require.NotNil(t, rt.Close())
}

type badStateChange struct{}

func (*badStateChange) Type() string { return "unknown" }

func TestRuntimeErrorsOnUnknownStateChanges(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	d := NewMockDriver(ctrl)
	p := NewMockPositionScaler(ctrl)
	s := NewMockStateMachine(ctrl)
	rt := &Runtime{
		Driver:         d,
		PositionScaler: p,
		StateMachine:   s,
	}
	s.EXPECT().Next().Return(true)
	s.EXPECT().Current().Return(&badStateChange{})
	s.EXPECT().Close().Return(nil)
	require.False(t, rt.Next())
	require.NotNil(t, rt.Close())
}

func TestRuntimeHandlesMove(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	d := NewMockDriver(ctrl)
	p := NewMockPositionScaler(ctrl)
	s := NewMockStateMachine(ctrl)
	rt := &Runtime{
		Driver:         d,
		PositionScaler: p,
		StateMachine:   s,
	}
	evt := &StateChangeMove{X: 1, Y: 2}
	s.EXPECT().Next().Return(true)
	s.EXPECT().Current().Return(evt)
	p.EXPECT().ScalePosition(evt.X, evt.Y).Return(2, 3)
	d.EXPECT().MoveMouse(2, 3).Return(nil)
	s.EXPECT().Next().Return(true)
	s.EXPECT().Current().Return(evt)
	p.EXPECT().ScalePosition(evt.X, evt.Y).Return(2, 3)
	d.EXPECT().MoveMouse(2, 3).Return(fmt.Errorf("move failed"))
	s.EXPECT().Close().Return(nil)

	require.True(t, rt.Next())
	require.False(t, rt.Next())
	require.NotNil(t, rt.Close())
}

func TestRuntimeHandlesDrag(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	d := NewMockDriver(ctrl)
	p := NewMockPositionScaler(ctrl)
	s := NewMockStateMachine(ctrl)
	rt := &Runtime{
		Driver:         d,
		PositionScaler: p,
		StateMachine:   s,
	}
	evt := &StateChangeDrag{X: 1, Y: 2}
	s.EXPECT().Next().Return(true)
	s.EXPECT().Current().Return(evt)
	p.EXPECT().ScalePosition(evt.X, evt.Y).Return(2, 3)
	d.EXPECT().DragMouse(2, 3).Return(nil)
	s.EXPECT().Next().Return(true)
	s.EXPECT().Current().Return(evt)
	p.EXPECT().ScalePosition(evt.X, evt.Y).Return(2, 3)
	d.EXPECT().DragMouse(2, 3).Return(fmt.Errorf("drag failed"))
	s.EXPECT().Close().Return(nil)

	require.True(t, rt.Next())
	require.False(t, rt.Next())
	require.NotNil(t, rt.Close())
}

func TestRuntimeHandlesClick(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	d := NewMockDriver(ctrl)
	p := NewMockPositionScaler(ctrl)
	s := NewMockStateMachine(ctrl)
	rt := &Runtime{
		Driver:         d,
		PositionScaler: p,
		StateMachine:   s,
	}
	evt := &StateChangeClick{}
	s.EXPECT().Next().Return(true)
	s.EXPECT().Current().Return(evt)
	d.EXPECT().Click().Return(nil)
	s.EXPECT().Next().Return(true)
	s.EXPECT().Current().Return(evt)
	d.EXPECT().Click().Return(fmt.Errorf("click failed"))
	s.EXPECT().Close().Return(nil)

	require.True(t, rt.Next())
	require.False(t, rt.Next())
	require.NotNil(t, rt.Close())
}

func TestRuntimeHandlesUnclick(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	d := NewMockDriver(ctrl)
	p := NewMockPositionScaler(ctrl)
	s := NewMockStateMachine(ctrl)
	rt := &Runtime{
		Driver:         d,
		PositionScaler: p,
		StateMachine:   s,
	}
	evt := &StateChangeUnclick{}
	s.EXPECT().Next().Return(true)
	s.EXPECT().Current().Return(evt)
	d.EXPECT().Unclick().Return(nil)
	s.EXPECT().Next().Return(true)
	s.EXPECT().Current().Return(evt)
	d.EXPECT().Unclick().Return(fmt.Errorf("unclick failed"))
	s.EXPECT().Close().Return(nil)

	require.True(t, rt.Next())
	require.False(t, rt.Next())
	require.NotNil(t, rt.Close())
}
