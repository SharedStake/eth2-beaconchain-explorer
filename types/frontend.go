package types

import (
	"time"

	"github.com/pkg/errors"
)

type EventName string

const (
	ValidatorBalanceDecreasedEventName              EventName = "validator_balance_decreased"
	ValidatorMissedProposalEventName                EventName = "validator_missed_proposal"
	ValidatorMissedAttestationEventName             EventName = "validator_missed_attestation"
	ValidatorGotSlashedEventName                    EventName = "validator_got_slashed"
	ValidatorDidSlashEventName                      EventName = "validator_did_slash"
	ValidatorStateChangedEventName                  EventName = "validator_state_changed"
	ValidatorReceivedDepositEventName               EventName = "validator_received_deposit"
	NetworkSlashingEventName                        EventName = "network_slashing"
	NetworkValidatorActivationQueueFullEventName    EventName = "network_validator_activation_queue_full"
	NetworkValidatorActivationQueueNotFullEventName EventName = "network_validator_activation_queue_not_full"
	NetworkValidatorExitQueueFullEventName          EventName = "network_validator_exit_queue_full"
	NetworkValidatorExitQueueNotFullEventName       EventName = "network_validator_exit_queue_not_full"
	NetworkLivenessIncreasedEventName               EventName = "network_liveness_increased"
)

var EventNames = []EventName{
	ValidatorBalanceDecreasedEventName,
	ValidatorMissedProposalEventName,
	ValidatorMissedAttestationEventName,
	ValidatorGotSlashedEventName,
	ValidatorDidSlashEventName,
	ValidatorStateChangedEventName,
	ValidatorReceivedDepositEventName,
	NetworkSlashingEventName,
	NetworkValidatorActivationQueueFullEventName,
	NetworkValidatorActivationQueueNotFullEventName,
	NetworkValidatorExitQueueFullEventName,
	NetworkValidatorExitQueueNotFullEventName,
	NetworkLivenessIncreasedEventName,
}

func EventNameFromString(event string) (EventName, error) {
	for _, en := range EventNames {
		if string(en) == event {
			return en, nil
		}
	}
	return "", errors.Errorf("Could not convert event to string. %v is not a known event type", event)
}

type Tag string

const (
	ValidatorTagsWatchlist Tag = "watchlist"
)

type Notification interface {
	GetSubscriptionID() uint64
	GetEventName() EventName
	GetEpoch() uint64
	GetInfo() string
}

type Subscription struct {
	ID           uint64     `db:"id"`
	UserID       uint64     `db:"user_id"`
	EventName    EventName  `db:"event_name"`
	EventFilter  string     `db:"event_filter"`
	LastSent     *time.Time `db:"last_sent_ts"`
	LastEpoch    *uint64    `db:"last_sent_epoch"`
	CreatedTime  time.Time  `db:"created_ts"`
	CreatedEpoch uint64     `db:"created_epoch"`
}

type TaggedValidators struct {
	UserID uint64 `db:"user_id"`
	Tag    string `db:"tag"`
	Validator
	Events []EventName `db:"events"`
}

type OAuthAppData struct {
	ID          uint64 `db:"id"`
	Owner       uint64 `db:"owner_id"`
	AppName     string `db:"app_name"`
	RedirectURI string `db:"redirect_uri"`
	Active      bool   `db:"active"`
}
