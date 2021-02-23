package permits

import "github.com/m3db/m3/src/x/context"

type fixedPermits struct {
	permits chan struct{}
}

var _ Permits = &fixedPermits{}

func NewFixedPermits(size int) Permits {
	permits := make(chan struct{}, size)
	for i := 0; i < size; i++ {
		permits <- struct{}{}
	}
	return &fixedPermits{permits: permits}
}

func (f *fixedPermits) Acquire(ctx context.Context) error {
	select {
	case <-f.permits:
		return nil
	case <-ctx.GoContext().Done():
		return ctx.GoContext().Err()
	}
}

func (f *fixedPermits) TryAcquire(ctx context.Context) (bool, error) {
	select {
	case <-ctx.GoContext().Done():
		return false, ctx.GoContext().Err()
	default:
	}

	select {
	case <-f.permits:
		return true, nil
	default:
		return false, nil
	}
}

func (f *fixedPermits) Release() {
	select {
	case f.permits <- struct{}{}:
	default:
		panic("Released more permits than acquired")
	}
}
